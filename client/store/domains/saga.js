import { name as querySliceName } from '@store/query/slice';
import { call, put, takeLatest } from 'redux-saga/effects';
import { search } from './api';
import { fetchDataSuccess } from './slice';

function* loadData({ payload: query }) {
  if (query.length === 0) {
    yield put(fetchDataSuccess([]));
    return;
  }

  const { data } = yield call(search, query);
  yield put(fetchDataSuccess(data));
}

export default function* searchSagas() {
  yield takeLatest(`${querySliceName}/update`, loadData);
}
