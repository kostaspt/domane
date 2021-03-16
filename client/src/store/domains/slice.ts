import { RootState } from '@/store';
import { createEntityAdapter, createSlice, EntityState } from '@reduxjs/toolkit';
import { STATUS_IDLE, STATUS_LOADING, STATUS_SUCCESS } from '@/store/status';

export const name: string = 'domains';

export type Domain = {
  domain: string;
  kind: string;
  position: number;
};

type State = EntityState<Domain> & {
  status: string;
};

export const domainIdSelector = (entity: Domain) => `${entity.position}-${entity.domain}`;

const adapter = createEntityAdapter<Domain>({
  selectId: domainIdSelector,
});

const slice = createSlice({
  name,
  initialState: adapter.getInitialState({
    status: STATUS_IDLE,
  }) as State,
  reducers: {
    fetchData(state) {
      state.status = STATUS_LOADING;
    },
    fetchDataSuccess(state, { payload }) {
      state.status = STATUS_SUCCESS;
      adapter.addMany(state, payload);
    },
  },
});

export const selectors = adapter.getSelectors((state: RootState) => state[name]);

export const { fetchDataSuccess } = slice.actions;
export default slice.reducer;
