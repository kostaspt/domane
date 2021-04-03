import {
  fetchDataFailure as domainFetchDataFailure,
  fetchDataSuccess as domainFetchDataSuccess,
} from '@store/domains/slice';
import { name as querySliceName } from '@store/query/slice';
import { call, put, select, takeLatest } from 'redux-saga/effects';
import { apiRequest } from './api';
import { fetchDataSuccess, fetchDataFailure, name } from './slice';

function* loadData() {
  const query: string = yield select((s) => s[querySliceName].text ?? '');

  if (query.length === 0) {
    yield put(fetchDataSuccess([]));
    return;
  }

  try {
    const { data } = yield call(apiRequest, query);
    yield put(fetchDataSuccess(data));
    yield put(domainFetchDataSuccess(data));
  } catch (err) {
    yield put(fetchDataFailure());
    yield put(domainFetchDataFailure());
  }
}

export default function* rootSaga() {
  yield takeLatest(`${name}/fetchData`, loadData);
}
