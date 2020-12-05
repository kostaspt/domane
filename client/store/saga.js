import { all } from 'redux-saga/effects';
import searchSagas from './search/saga';

export default function* rootSaga() {
  yield all([searchSagas()]);
}
