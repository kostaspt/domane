import { all } from 'redux-saga/effects';
import domainsSagas from './domains/saga';

export default function* rootSaga() {
  yield all([domainsSagas()]);
}
