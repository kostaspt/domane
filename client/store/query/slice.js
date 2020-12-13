import { createSlice } from '@reduxjs/toolkit';

export const name = 'query';

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

const { actions, reducer } = slice;
export const { update } = actions;
export default reducer;
