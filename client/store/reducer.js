import { combineReducers } from '@reduxjs/toolkit';
import domainsReducer, { name as domainsSliceName } from '@store/domains/slice';
import queryReducer, { name as querySliceName } from '@store/query/slice';

export default combineReducers({
  [domainsSliceName]: domainsReducer,
  [querySliceName]: queryReducer,
});
