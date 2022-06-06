import { styled, Paper, createTheme } from '@mui/material';

const theme = createTheme();

const BasePaper = styled(Paper)({
  position: 'relative',
  padding: theme.spacing(2),
  display: 'flex',
  overflow: 'auto',
  flexDirection: 'row',
  height: 400,
});

export { BasePaper };
