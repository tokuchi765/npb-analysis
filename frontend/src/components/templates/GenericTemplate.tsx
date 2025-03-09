import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import {
  Accordion,
  AccordionDetails,
  AccordionSummary,
  AppBar,
  Box,
  Button,
  CssBaseline,
  Drawer,
  Grid2,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Paper,
  Toolbar,
  Typography,
} from '@mui/material';
import {
  Pentagon,
  Search,
  SportsBaseball,
  TableChart,
  Person,
  Home,
  Group,
  SportsCricket,
  ExpandMore,
} from '@mui/icons-material';
import Title from '../common/Title';

const drawerWidth = 300;

const Copyright = () => {
  return (
    <React.Fragment>
      <Grid2
        container
        paddingBottom={3}
        alignItems="center"
        sx={{ display: 'flex', justifyContent: 'center' }}
      >
        <Typography marginBottom={3} variant="body2" color="textSecondary" align="center">
          {'Copyright ? '}
          <Link color="inherit" to="/">
            プロ野球データ分析
          </Link>{' '}
          {new Date().getFullYear()}
          {'.'}
        </Typography>
      </Grid2>
      <Grid2 container sx={{ display: 'flex' }}>
        <Grid2 width={'50%'}>
          <Title>参考サイト</Title>
          <a href="https://npb.jp/" target="_blank" rel="noopener noreferrer">
            日本野球機構公式サイト
          </a>
        </Grid2>
        <Grid2 width={'50%'}>
          <Title>免責事項</Title>
          <Typography sx={{ fontSize: '12px' }} color="textSecondary">
            このサイトはNPB公式サイトより集計したデータを独自に再計算し提示しているサイトです。
            <br />
            このサイトにより生じた損害に付きましては一切責任を負いません。
            <br />
            自己責任でご使用頂くようお願い申し上げます。
          </Typography>
        </Grid2>
      </Grid2>
    </React.Fragment>
  );
};

export interface GenericTemplateProps {
  children: React.ReactNode;
  title: string;
  displayBackButton?: boolean;
  backOnClick?: () => void;
}

function GenericTemplate(props: GenericTemplateProps) {
  const [expanded, setExpanded] = useState(() => {
    const accordionExpanded = localStorage.getItem('accordionExpanded');
    if (accordionExpanded) {
      return JSON.parse(accordionExpanded) || false;
    }
    return false;
  });
  const handleExpandedChange = () => {
    setExpanded((prevExpanded: any) => {
      const newExpanded = !prevExpanded;
      localStorage.setItem('accordionExpanded', JSON.stringify(newExpanded));
      return newExpanded;
    });
  };

  return (
    <Box sx={{ display: 'flex' }}>
      <CssBaseline />
      {/* AppBar（ヘッダー） */}
      <AppBar sx={{ zIndex: (theme) => theme.zIndex.drawer + 1 }}>
        <Toolbar>
          <Typography variant="h6">プロ野球データ分析</Typography>
        </Toolbar>
      </AppBar>
      {/* Drawer（サイドメニュー） */}
      <Drawer
        variant="permanent"
        anchor="left"
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': { width: drawerWidth, boxSizing: 'border-box' },
        }}
      >
        <Toolbar />
        <List>
          <ListItem component={Link} to="/">
            <ListItemIcon>
              <Home />
            </ListItemIcon>
            <ListItemText primary="トップページ" />
          </ListItem>
          <Accordion expanded={expanded} onChange={handleExpandedChange} disableGutters>
            <AccordionSummary
              style={{
                justifyContent: 'flex-start',
                display: 'flex',
                alignItems: 'center',
                margin: 0,
                padding: 0,
                marginLeft: 17,
              }}
              expandIcon={<ExpandMore />}
            >
              <ListItemIcon>
                <Group />
              </ListItemIcon>
              <ListItemText primary="チーム情報" />
            </AccordionSummary>
            <AccordionDetails>
              <ListItem component={Link} to="/season">
                <ListItemIcon>
                  <TableChart />
                </ListItemIcon>
                <ListItemText primary="シーズン成績ページ" />
              </ListItem>
              <ListItem component={Link} to="/batting">
                <ListItemIcon>
                  <SportsCricket />
                </ListItemIcon>
                <ListItemText primary="打撃成績ページ" />
              </ListItem>
              <ListItem component={Link} to="/pitching">
                <ListItemIcon>
                  <SportsBaseball />
                </ListItemIcon>
                <ListItemText primary="投手成績ページ" />
              </ListItem>
              <ListItem component={Link} to="/strength">
                <ListItemIcon>
                  <Pentagon />
                </ListItemIcon>
                <ListItemText primary="チーム戦力チャート" />
              </ListItem>
            </AccordionDetails>
          </Accordion>
          <ListItem component={Link} to="/players">
            <ListItemIcon>
              <Person />
            </ListItemIcon>
            <ListItemText primary="選手一覧ページ" />
          </ListItem>
          <ListItem component={Link} to="/manager">
            <ListItemIcon>
              <Person />
            </ListItemIcon>
            <ListItemText primary="監督ページ" />
          </ListItem>
          <ListItem component={Link} to="/search">
            <ListItemIcon>
              <Search />
            </ListItemIcon>
            <ListItemText primary="選手検索ページ" />
          </ListItem>
          <ListItem component={Link} to="/search_grades">
            <ListItemIcon>
              <Search />
            </ListItemIcon>
            <ListItemText primary="選手成績検索ページ" />
          </ListItem>
        </List>
      </Drawer>
      <Paper
        elevation={3}
        sx={{
          marginTop: '64px',
          height: '100%',
          width: '100%',
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'space-between',
          p: 2,
        }}
      >
        {/* メインコンテンツ */}
        <Box component="main" sx={{ flexGrow: 1, p: 3, transition: 'margin 0.3s' }}>
          <Grid2 sx={{ display: 'flex' }}>
            <Typography variant="h5">{props.title}</Typography>
            {props.displayBackButton && (
              <Button
                sx={{ marginLeft: 2 }}
                onClick={props.backOnClick}
                variant="contained"
                color="primary"
              >
                戻る
              </Button>
            )}
          </Grid2>
        </Box>

        {props.children}
        <Box component="footer" sx={{ borderTop: '1px solid #ccc', pt: 1 }}>
          <Copyright />
        </Box>
      </Paper>
    </Box>
  );
}

export default GenericTemplate;
