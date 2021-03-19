import { createEntityAdapter, createSlice } from '@reduxjs/toolkit';
import { RootState } from '@store';

export const name: string = 'availabilities';

export type Availability = {
  domain: string;
  isAvailable: boolean;
};

const adapter = createEntityAdapter<Availability>({
  selectId: (entity) => entity.domain,
});

const slice = createSlice({
  name,
  initialState: adapter.getInitialState(),
  reducers: {
    upsert(state, { payload }) {
      adapter.upsertOne(state, payload);
    },
  },
});

export const selectors = adapter.getSelectors((state: RootState) => state[name]);

export const { upsert } = slice.actions;
export default slice.reducer;
