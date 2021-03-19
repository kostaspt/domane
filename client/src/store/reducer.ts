import { combineReducers } from '@reduxjs/toolkit';
import availabilitiesReducer, { name as availabilitiesSliceName } from '@store/availabilities/slice';
import domainsReducer, { name as domainsSliceName } from '@store/domains/slice';
import domainsExtensionsReducer, { name as domainsExtensionsSliceName } from '@store/domainsExtensions/slice';
import domainsSimilarsReducer, { name as domainsSimilarsSliceName } from '@store/domainsSimilars/slice';
import queryReducer, { name as querySliceName } from '@store/query/slice';

export default combineReducers({
  [availabilitiesSliceName]: availabilitiesReducer,
  [domainsExtensionsSliceName]: domainsExtensionsReducer,
  [domainsSimilarsSliceName]: domainsSimilarsReducer,
  [domainsSliceName]: domainsReducer,
  [querySliceName]: queryReducer,
});
