import { fetchDataSuccess as fetchDomainsSuccess } from '@store/domains/slice';
import { name as querySliceName } from '@store/query/slice';
import { call, put, select, takeLatest } from 'redux-saga/effects';
import { apiRequest } from './api';
import { name, fetchDataSuccess } from './slice';

function* loadData() {
  const query = yield select((s) => s[querySliceName].text ?? '');

  if (query.length === 0) {
    yield put(fetchDataSuccess([]));
    return;
  }

  const { data } = yield call(apiRequest, query);
  yield put(fetchDataSuccess(data));
  yield put(fetchDomainsSuccess(data));
}

export default function* rootSaga() {
  yield takeLatest(`${name}/fetchData`, loadData);
}
