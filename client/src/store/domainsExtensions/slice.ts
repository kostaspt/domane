import { RootState } from '@store';
import { createSelector, createSlice, EntityId } from '@reduxjs/toolkit';
import { domainIdSelector } from '@store/domains/slice';
import { STATUS_FAILURE, STATUS_IDLE, STATUS_LOADING, STATUS_SUCCESS } from '@store/status';

export const name: string = 'domainsExtensions';

type State = {
  ids: EntityId[];
  status: string;
};

const slice = createSlice({
  name,
  initialState: {
    ids: [],
    status: STATUS_IDLE,
  } as State,
  reducers: {
    fetchData(state) {
      state.status = STATUS_LOADING;
    },
    fetchDataSuccess(state, { payload }) {
      state.status = STATUS_SUCCESS;
      state.ids = payload.map(domainIdSelector);
    },
    fetchDataFailure(state) {
      state.status = STATUS_FAILURE;
    },
  },
});

export const selectIds = createSelector(
  (state: RootState) => state[name],
  (s: State) => s.ids
);

export const selectStatus = createSelector(
  (state: RootState) => state[name],
  (s: State) => s.status
);

export const { fetchData, fetchDataSuccess, fetchDataFailure } = slice.actions;
export default slice.reducer;
