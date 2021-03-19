import { fetchData as domainsExtensionsFetchData } from '@store/domainsExtensions/slice';
import { fetchData as domainsSimilarsFetchData } from '@store/domainsSimilars/slice';
import { put, takeLatest } from 'redux-saga/effects';
import { name } from './slice';

function* updateQuery() {
  yield put(domainsExtensionsFetchData());
  yield put(domainsSimilarsFetchData());
}

export default function* rootSaga() {
  yield takeLatest(`${name}/update`, updateQuery);
}
