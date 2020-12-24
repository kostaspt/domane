import { all } from 'redux-saga/effects';
import availabilitiesSagas from './availabilities/saga';
import domainsSagas from './domains/saga';
import querySagas from './query/saga';

export default function* rootSaga() {
  yield all([availabilitiesSagas(), domainsSagas(), querySagas()]);
}
