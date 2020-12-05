import { createEntityAdapter, createSlice } from '@reduxjs/toolkit';

const adapter = createEntityAdapter({
  selectId: (result) => result.position,
});

export const name = 'search'

const initialState = {
  ids: [],
  entities: {},
  query: '',
  loading: false,
}

const searchSlice = createSlice({
  name,
  initialState,
  reducers: {
    fetchDataSuccess(state, { payload }) {
      state.loading = false;
      adapter.setAll(state, payload);
    },
    updateQuery(state, { payload }) {
      state.loading = true;
      state.query = payload;
    },
  },
});

export const searchSelector = adapter.getSelectors((state) => state[name])

const { actions, reducer } = searchSlice;
export const { fetchDataSuccess, updateQuery } = actions;
export default reducer;
