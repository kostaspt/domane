import { querySliceName } from '@store/query/slice';
import { call, put, select, takeLatest } from 'redux-saga/effects';
import { search } from './api';
import { domainsSliceName, fetchDomainsSuccess } from './slice';

function* loadData() {
  const query = yield select((s) => s[querySliceName].text ?? '');

  if (query.length === 0) {
    yield put(fetchDomainsSuccess([]));
    return;
  }

  const { data } = yield call(search, query);
  yield put(fetchDomainsSuccess(data));
}

export default function* rootSaga() {
  yield takeLatest(`${domainsSliceName}/fetchData`, loadData);
}
