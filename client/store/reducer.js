import searchReducer from '@store/search/slice';
import { combineReducers } from '@reduxjs/toolkit';

export default combineReducers({
  search: searchReducer,
});
