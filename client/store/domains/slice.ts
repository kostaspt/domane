import { createEntityAdapter, createSlice } from '@reduxjs/toolkit';
import { RootState } from '@/store';

const name: string = 'domains';

type Domain = {
  domain: string;
  position: number;
};

const adapter = createEntityAdapter<Domain>({
  selectId: (entity) => `${entity.position}-${entity.domain}`,
});

const slice = createSlice({
  name,
  initialState: adapter.getInitialState({
    loading: false,
  }),
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

export const domainsSelector = adapter.getSelectors((state: RootState) => state[name]);

export const domainsSliceName = name;
export const { fetchData: fetchDomains, fetchDataSuccess: fetchDomainsSuccess } = slice.actions;
export default slice.reducer;
