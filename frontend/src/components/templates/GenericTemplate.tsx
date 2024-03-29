import React from 'react';
import clsx from 'clsx';
import { createTheme } from '@material-ui/core/styles';
import * as colors from '@material-ui/core/colors';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import { ThemeProvider } from '@material-ui/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Drawer from '@material-ui/core/Drawer';
import Box from '@material-ui/core/Box';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import List from '@material-ui/core/List';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import Container from '@material-ui/core/Container';
import { Link } from 'react-router-dom';
import IconButton from '@material-ui/core/IconButton';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import { Menu } from '@mui/material';
import {
  Groups,
  SportsCricket,
  SportsBaseball,
  Person,
  Home,
  TableChart,
  ChevronLeft,
  Pentagon,
} from '@mui/icons-material';
import MenuIcon from '@mui/icons-material/Menu';

const drawerWidth = 240;

const theme = createTheme({
  typography: {
    fontFamily: [
      'Noto Sans JP',
      'Lato',
      '游ゴシック Medium',
      '游ゴシック体',
      'Yu Gothic Medium',
      'YuGothic',
      'ヒラギノ角ゴ ProN',
      'Hiragino Kaku Gothic ProN',
      'メイリオ',
      'Meiryo',
      'ＭＳ Ｐゴシック',
      'MS PGothic',
      'sans-serif',
    ].join(','),
  },
  palette: {
    primary: { main: colors.blue[800] }, // テーマの色
  },
});

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
    },
    toolbar: {
      paddingRight: 24,
    },
    toolbarIcon: {
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'flex-end',
      padding: '0 8px',
      ...theme.mixins.toolbar,
    },
    appBar: {
      zIndex: theme.zIndex.drawer + 1,
      transition: theme.transitions.create(['width', 'margin'], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
    },
    appBarShift: {
      marginLeft: drawerWidth,
      width: `calc(100% - ${drawerWidth}px)`,
      transition: theme.transitions.create(['width', 'margin'], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    menuButton: {
      marginRight: 36,
    },
    menuButtonHidden: {
      display: 'none',
    },
    title: {
      flexGrow: 1,
    },
    pageTitle: {
      marginBottom: theme.spacing(1),
    },
    drawerPaper: {
      position: 'relative',
      whiteSpace: 'nowrap',
      width: drawerWidth,
      transition: theme.transitions.create('width', {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    drawerPaperClose: {
      overflowX: 'hidden',
      transition: theme.transitions.create('width', {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
      width: theme.spacing(7),
      [theme.breakpoints.up('sm')]: {
        width: theme.spacing(9),
      },
    },
    appBarSpacer: theme.mixins.toolbar,
    content: {
      flexGrow: 1,
      height: '100vh',
      overflow: 'auto',
    },
    container: {
      paddingTop: theme.spacing(4),
      paddingBottom: theme.spacing(4),
    },
    paper: {
      padding: theme.spacing(2),
      display: 'flex',
      overflow: 'auto',
      flexDirection: 'column',
    },
    link: {
      textDecoration: 'none',
      color: theme.palette.text.secondary,
    },
  })
);

const Copyright = () => {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" to="/">
        管理画面
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
};

export interface GenericTemplateProps {
  children: React.ReactNode;
  title: string;
}

function GenericTemplate(props: GenericTemplateProps) {
  const classes = useStyles();
  const [open, setOpen] = React.useState(true);
  const handleDrawerOpen = () => {
    setOpen(true);
  };
  const handleDrawerClose = () => {
    setOpen(false);
  };
  const [anchorEl, setAnchorEl] = React.useState(null);

  function handleClick(event: any) {
    if (anchorEl !== event.currentTarget) {
      setAnchorEl(event.currentTarget);
    }
  }

  function handleClose() {
    setAnchorEl(null);
  }

  return (
    <ThemeProvider theme={theme}>
      <div className={classes.root}>
        <CssBaseline />
        <AppBar position="absolute" className={clsx(classes.appBar, open && classes.appBarShift)}>
          <Toolbar className={classes.toolbar}>
            <IconButton
              edge="start"
              color="inherit"
              aria-label="open drawer"
              onClick={handleDrawerOpen}
              className={clsx(classes.menuButton, open && classes.menuButtonHidden)}
            >
              <MenuIcon />
            </IconButton>
            <Typography
              component="h1"
              variant="h6"
              color="inherit"
              noWrap
              className={classes.title}
            >
              管理画面
            </Typography>
          </Toolbar>
        </AppBar>
        <Drawer
          variant="permanent"
          classes={{
            paper: clsx(classes.drawerPaper, !open && classes.drawerPaperClose),
          }}
          open={open}
        >
          <div className={classes.toolbarIcon}>
            <IconButton onClick={handleDrawerClose}>
              <ChevronLeft />
            </IconButton>
          </div>
          <Divider />
          <List>
            <Link to="/" className={classes.link}>
              <ListItem button>
                <ListItemIcon>
                  <Home />
                </ListItemIcon>
                <ListItemText primary="トップページ" />
              </ListItem>
            </Link>
            <ListItem className={classes.link} onMouseOver={handleClick}>
              <ListItemIcon>
                <Groups />
              </ListItemIcon>
              <ListItemText primary="チーム情報" />
              <Menu
                anchorEl={anchorEl}
                anchorOrigin={{ vertical: 'top', horizontal: 'right' }}
                open={Boolean(anchorEl)}
                onClose={handleClose}
                MenuListProps={{ onMouseLeave: handleClose }}
              >
                <Link to="/season" className={classes.link}>
                  <ListItem button>
                    <ListItemIcon>
                      <TableChart />
                    </ListItemIcon>
                    <ListItemText primary="シーズン成績ページ" />
                  </ListItem>
                </Link>
                <Link to="/batting" className={classes.link}>
                  <ListItem button>
                    <ListItemIcon>
                      <SportsCricket />
                    </ListItemIcon>
                    <ListItemText primary="打撃成績ページ" />
                  </ListItem>
                </Link>
                <Link to="/pitching" className={classes.link}>
                  <ListItem button>
                    <ListItemIcon>
                      <SportsBaseball />
                    </ListItemIcon>
                    <ListItemText primary="投手成績ページ" />
                  </ListItem>
                </Link>
                <Link to="/strength" className={classes.link}>
                  <ListItem button>
                    <ListItemIcon>
                      <Pentagon />
                    </ListItemIcon>
                    <ListItemText primary="チーム戦力チャート" />
                  </ListItem>
                </Link>
              </Menu>
            </ListItem>
            <Link to="/players" className={classes.link}>
              <ListItem button>
                <ListItemIcon>
                  <Person />
                </ListItemIcon>
                <ListItemText primary="選手一覧ページ" />
              </ListItem>
            </Link>
            <Link to="/manager" className={classes.link}>
              <ListItem button>
                <ListItemIcon>
                  <Person />
                </ListItemIcon>
                <ListItemText primary="監督ページ" />
              </ListItem>
            </Link>
          </List>
        </Drawer>
        <main className={classes.content}>
          <div className={classes.appBarSpacer} />
          <Container maxWidth="lg" className={classes.container}>
            <Typography
              component="h2"
              variant="h5"
              color="inherit"
              noWrap
              className={classes.pageTitle}
            >
              {props.title}
            </Typography>
            {props.children}
            <Box pt={4}>
              <Copyright />
            </Box>
          </Container>
        </main>
      </div>
    </ThemeProvider>
  );
}

export default GenericTemplate;
