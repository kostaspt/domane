import { createEntityAdapter, createSlice } from '@reduxjs/toolkit';
import { RootState } from '@/store';

const name: string = 'availabilities';

type Availability = {
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

export const availabilitiesSelector = adapter.getSelectors((state: RootState) => state[name]);

export const availabilitiesSliceName = name;
export const { upsert: upsertAvailability } = slice.actions;
export default slice.reducer;
