import { createSlice } from '@reduxjs/toolkit';

const name = 'query';

const initialState = {
  text: '',
};

const slice = createSlice({
  name,
  initialState,
  reducers: {
    update(state, { payload }) {
      state.text = payload;
    },
  },
});

export const querySliceName = name;
export const { update: updateQuery } = slice.actions;
export default slice.reducer;
