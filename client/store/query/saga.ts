import { fetchDomains } from '@store/domains/slice';
import { put, takeLatest } from 'redux-saga/effects';
import { querySliceName } from './slice';

function* updateQuery() {
  yield put(fetchDomains());
}

export default function* rootSaga() {
  yield takeLatest(`${querySliceName}/update`, updateQuery);
}
