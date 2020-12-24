import { createEntityAdapter, createSlice } from '@reduxjs/toolkit';

const adapter = createEntityAdapter({
  selectId: (entity) => `${entity.position}-${entity.domain}`,
});

export const name = 'domains';

const initialState = {
  ids: [],
  entities: {},
  loading: false,
};

const slice = createSlice({
  name,
  initialState,
  reducers: {
    fetchData(state) {
      state.loading = true;
    },
    fetchDataSuccess(state, { payload }) {
      state.loading = false;
      adapter.setAll(state, payload);
    },
  },
});

export const domainsSelector = adapter.getSelectors((state) => state[name]);

const { actions, reducer } = slice;
export const { fetchData, fetchDataSuccess } = actions;
export default reducer;
