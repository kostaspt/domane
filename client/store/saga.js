import { all } from 'redux-saga/effects';
import domainsSagas from './domains/saga';
import querySagas from './query/saga';

export default function* rootSaga() {
  yield all([domainsSagas(), querySagas()]);
}
