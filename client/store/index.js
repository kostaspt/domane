import { configureStore, getDefaultMiddleware } from '@reduxjs/toolkit';
import camelMiddleware from 'redux-camel';
import logger from 'redux-logger';
import createSagaMiddleware from 'redux-saga';
import rootReducer from './reducer';
import rootSaga from './saga';

const sagaMiddleware = createSagaMiddleware();
const middleware = [camelMiddleware(), ...getDefaultMiddleware({ thunk: false }), sagaMiddleware];

// const devMode = process.env.NODE_ENV !== 'production';
const devMode = true;
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
