import { name as querySliceName } from '@store/query/slice';
import { call, put, select, takeLatest } from 'redux-saga/effects';
import { search } from './api';
import { fetchDataSuccess, name as domainsSliceName } from './slice';

function* loadData() {
  const query = yield select((s) => s[querySliceName].text ?? '');

  if (query.length === 0) {
    yield put(fetchDataSuccess([]));
    return;
  }

  const { data } = yield call(search, query);
  yield put(fetchDataSuccess(data));
}

export default function* rootSaga() {
  yield takeLatest(`${domainsSliceName}/fetchData`, loadData);
}
