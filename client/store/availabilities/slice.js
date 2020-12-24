import { createEntityAdapter, createSlice } from '@reduxjs/toolkit';

const adapter = createEntityAdapter({
  selectId: (entity) => entity.domain,
});

const name = 'availabilities';

const initialState = {
  ids: [],
  entities: {},
};

const slice = createSlice({
  name,
  initialState,
  reducers: {
    upsert(state, { payload }) {
      adapter.upsertOne(state, payload);
    },
  },
});

export const availabilitiesSelector = adapter.getSelectors((state) => state[name]);

export const availabilitiesSliceName = name;
export const { upsert: upsertAvailability } = slice.actions;
export default slice.reducer;
