import { configureStore, getDefaultMiddleware } from '@reduxjs/toolkit';
import createSagaMiddleware from 'redux-saga';
import logger from 'redux-logger';
import rootReducer from './reducer';
import rootSaga from './saga';

const sagaMiddleware = createSagaMiddleware();
const middleware = [...getDefaultMiddleware({ thunk: false }), sagaMiddleware];

const devMode = process.env.NODE_ENV !== 'production';
if (devMode) {
  middleware.push(logger);
}

const store = configureStore({
  reducer: rootReducer,
  middleware,
  devTools: devMode,
});

sagaMiddleware.run(rootSaga);

export default store;
