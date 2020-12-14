import { fetchData as fetchDomainData } from '@store/domains/slice';
import { put, takeLatest } from 'redux-saga/effects';
import { name as querySliceName } from './slice';

function* updateQuery() {
  yield put(fetchDomainData());
}

export default function* rootSaga() {
  yield takeLatest(`${querySliceName}/update`, updateQuery);
}
