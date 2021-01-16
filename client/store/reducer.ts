import { combineReducers } from '@reduxjs/toolkit';
import availabilitiesReducer, { availabilitiesSliceName } from '@store/availabilities/slice';
import domainsReducer, { domainsSliceName } from '@store/domains/slice';
import queryReducer, { querySliceName } from '@store/query/slice';

export default combineReducers({
  [availabilitiesSliceName]: availabilitiesReducer,
  [domainsSliceName]: domainsReducer,
  [querySliceName]: queryReducer,
});
