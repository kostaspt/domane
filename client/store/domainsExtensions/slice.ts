import { RootState } from '@/store';
import { createSelector, createSlice, EntityId } from '@reduxjs/toolkit';
import { domainIdSelector } from '@store/domains/slice';
import { STATUS_IDLE, STATUS_LOADING, STATUS_SUCCESS } from '@store/status';

export const name: string = 'domainsExtensions';

type State = {
  ids: EntityId[];
  status: string;
};

const initialState = {
  ids: [],
  status: STATUS_IDLE,
} as State;

const slice = createSlice({
  name,
  initialState,
  reducers: {
    fetchData(state) {
      state.status = STATUS_LOADING;
    },
    fetchDataSuccess(state, { payload }) {
      state.status = STATUS_SUCCESS;
      state.ids = payload.map(domainIdSelector);
    },
  },
});

export const selectIds = createSelector(
  (state: RootState) => state[name],
  (s: State) => s.ids
);

export const { fetchData, fetchDataSuccess } = slice.actions;
export default slice.reducer;
