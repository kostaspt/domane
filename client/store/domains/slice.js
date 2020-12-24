import { createEntityAdapter, createSlice } from '@reduxjs/toolkit';

const adapter = createEntityAdapter({
  selectId: (entity) => `${entity.position}-${entity.domain}`,
});

const name = 'domains';

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

export const domainsSliceName = name;
export const { fetchData: fetchDomains, fetchDataSuccess: fetchDomainsSuccess } = slice.actions;
export default slice.reducer;
