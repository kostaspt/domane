import { Dictionary, EntityId } from '@reduxjs/toolkit';
import { Domain, name as domainsSliceName, selectors as domainsSelectors } from '@store/domains/slice';
import isNode from 'detect-node';
import getConfig from 'next/config';
import ReconnectingWebSocket from 'reconnecting-websocket';
import { EventChannel, eventChannel } from 'redux-saga';
import { call, delay, fork, put, select, take, takeLatest } from 'redux-saga/effects';
import { Availability, selectors, upsert } from './slice';

const { publicRuntimeConfig } = getConfig();
const { API_URL } = publicRuntimeConfig;

let ws: ReconnectingWebSocket;

function createEventChannel() {
  return eventChannel((emit) => {
    ws.addEventListener('message', (lastMessage) => {
      try {
        const data = JSON.parse(lastMessage?.data);
        if (data?.domain) {
          return emit({ data });
        }
      } catch (e) {
        console.debug(e);
      }
    });

    return () => ws.close();
  });
}

function* setupConnection() {
  if (isNode) return;

  const wsURL = API_URL.replace('http', 'ws');
  ws = new ReconnectingWebSocket(`${wsURL}/ws/whois`);

  const channel: EventChannel<any> = yield call(createEventChannel);
  while (true) {
    const { data } = yield take(channel);
    yield put(upsert(data));
  }
}

function* sendMessages() {
  yield delay(500);

  const ids: EntityId[] = yield select(domainsSelectors.selectIds);
  const domains: Dictionary<Domain> = yield select(domainsSelectors.selectEntities);
  const availabilities: Dictionary<Availability> = yield select(selectors.selectEntities);

  ids
    .filter((id) => typeof availabilities[domains[id]?.domain ?? '']?.isAvailable === 'undefined')
    .forEach((id) => ws.send(domains[id]?.domain ?? ''));
}

export default function* rootSaga() {
  yield fork(setupConnection);
  yield takeLatest(`${domainsSliceName}/fetchDataSuccess`, sendMessages);
}
