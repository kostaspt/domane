import { all } from 'redux-saga/effects';
import availabilitiesSagas from './availabilities/saga';
import domainsExtensionsSagas from './domainsExtensions/saga';
import domainsSimilarsSagas from './domainsSimilars/saga';
import querySagas from './query/saga';

export default function* rootSaga() {
  yield all([availabilitiesSagas(), domainsExtensionsSagas(), domainsSimilarsSagas(), querySagas()]);
}
