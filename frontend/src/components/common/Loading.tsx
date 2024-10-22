import React from 'react';
import { CircularProgress, styled } from '@mui/material';

const LoadingDiv = styled('div')({
  position: 'fixed',
  top: 0,
  left: 0,
  width: '100%',
  height: '100%',
  backgroundColor: 'rgba(0, 0, 0, 0.5)',
  zIndex: 9999,
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
});

function Loading(props: { loading: boolean }) {
  return (
    <React.Fragment>
      {props.loading && (
        <LoadingDiv>
          <CircularProgress size="10%" />
        </LoadingDiv>
      )}
    </React.Fragment>
  );
}

export default Loading;
