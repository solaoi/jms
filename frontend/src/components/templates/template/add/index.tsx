import React, { FC, useState } from "react";
import clsx from "clsx";
import { useDrop } from "react-dnd";
import {
  AppBar,
  Badge,
  Box,
  Container,
  Divider,
  Drawer,
  CssBaseline,
  FormControl,
  FormGroup,
  IconButton,
  List,
  Switch,
  Toolbar,
  Typography,
  Card,
  CardHeader,
  CardContent,
  TextField,
  Grid,
} from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";
import {
  Menu,
  ChevronLeft,
  Create,
  Notifications,
  EditAttributes,
  LooksOne,
} from "@material-ui/icons";
import Copyright from "~/components/organizms/Copyright";
import NestedIcon from "~/components/atoms/NestedIcon";
import TemplateElementList, {
  DRAG_EVENT_TYPE,
} from "~/components/templates/template/add/TemplateElementList";
import {
  TemplateElementType,
  TEMPLATE_ELEMENT,
} from "~/components/templates/template/add/TemplateElement";

const SYSTEM_NAME = "JsonManagementSystem";
const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
  root: {
    display: "flex",
  },
  toolbar: {
    paddingRight: 24, // keep right padding when drawer closed
  },
  toolbarIcon: {
    display: "flex",
    alignItems: "center",
    justifyContent: "flex-end",
    padding: "0 8px",
    ...theme.mixins.toolbar,
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
    transition: theme.transitions.create(["width", "margin"], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },
  appBarShift: {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(["width", "margin"], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  menuButton: {
    marginRight: 36,
  },
  menuButtonHidden: {
    display: "none",
  },
  title: {
    flexGrow: 1,
  },
  drawerPaper: {
    position: "relative",
    whiteSpace: "nowrap",
    width: drawerWidth,
    transition: theme.transitions.create("width", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  drawerPaperClose: {
    overflowX: "hidden",
    transition: theme.transitions.create("width", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    width: theme.spacing(7),
    [theme.breakpoints.up("sm")]: {
      width: theme.spacing(9),
    },
  },
  appBarSpacer: theme.mixins.toolbar,
  content: {
    flexGrow: 1,
    height: "100vh",
    overflow: "auto",
  },
  container: {
    paddingTop: theme.spacing(4),
    paddingBottom: theme.spacing(4),
  },
  paper: {
    padding: theme.spacing(2),
    display: "flex",
    overflow: "auto",
    flexDirection: "column",
  },
  fixedHeight: {
    height: 240,
  },
}));

const Main: FC = () => {
  const classes = useStyles();
  const [open, setOpen] = useState(true);
  const [list, setList] = useState<string[]>([]);
  const [state, setState] = useState<{ [index: string]: string | boolean }>({});
  const handleStateChange = (val: { [index: string]: string | boolean }) => {
    setState((prev) => ({ ...prev, ...val }));
  };
  const handleDrawerOpen = () => {
    setOpen(true);
  };
  const handleDrawerClose = () => {
    setOpen(false);
  };

  const [, drop] = useDrop({
    accept: DRAG_EVENT_TYPE.TEMPLATE_ELEMENT,
    drop: (val: TemplateElementType) => {
      console.log(val);
      setList((prev) => [...prev, val.text]);
    },
  });

  return (
    <div className={classes.root}>
      <CssBaseline />
      <AppBar
        position="absolute"
        className={clsx(classes.appBar, open && classes.appBarShift)}
      >
        <Toolbar className={classes.toolbar}>
          <IconButton
            edge="start"
            color="inherit"
            aria-label="open drawer"
            onClick={handleDrawerOpen}
            className={clsx(
              classes.menuButton,
              open && classes.menuButtonHidden
            )}
          >
            <Menu />
          </IconButton>
          <Typography
            component="h1"
            variant="h6"
            color="inherit"
            noWrap
            className={classes.title}
          >
            {SYSTEM_NAME}
          </Typography>
          <IconButton color="inherit">
            <Badge badgeContent={4} color="secondary">
              <Notifications />
            </Badge>
          </IconButton>
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
          <TemplateElementList />
        </List>
        <Divider />
      </Drawer>
      <main className={classes.content}>
        <div className={classes.appBarSpacer} />
        <Container maxWidth="lg" className={classes.container}>
          <Card ref={drop} style={{ backgroundColor: "#EEE" }}>
            <CardHeader
              title="NewTemplate"
              subheader={
                <div>we should Drag and Drop the left side elements below.</div>
              }
            />
            <CardContent>
              <FormControl component="fieldset" fullWidth>
                <FormGroup aria-label="position">
                  {list.map((text, index) => {
                    if (text === TEMPLATE_ELEMENT.BOOLEAN) {
                      const listKey = `form-boolean-${index}`;
                      const stateKey = `boolean-${index}`;
                      return (
                        <Card key={listKey} style={{ margin: "8px" }}>
                          <CardHeader
                            avatar={<EditAttributes />}
                            title="Boolean"
                            titleTypographyProps={{ variant: "subtitle1" }}
                          />
                          <CardContent>
                            <TextField
                              placeholder="Title"
                              margin="normal"
                              style={{ marginLeft: "8px" }}
                            />
                            <div>
                              <Switch
                                checked={state[stateKey] ? true : false}
                                onChange={() => {
                                  handleStateChange({
                                    [stateKey]: !state[stateKey],
                                  });
                                }}
                                name="checkedB"
                                color="primary"
                              />
                            </div>
                          </CardContent>
                        </Card>
                      );
                    } else if (text === TEMPLATE_ELEMENT.STRING) {
                      const listKey = `form-string-${index}`;
                      return (
                        <Card key={listKey} style={{ margin: "8px" }}>
                          <CardHeader
                            avatar={<Create />}
                            title="String"
                            titleTypographyProps={{ variant: "subtitle1" }}
                          />
                          <CardContent>
                            <TextField
                              placeholder="Title"
                              margin="normal"
                              style={{ marginLeft: "8px" }}
                            />
                            <TextField
                              placeholder="DefaultValue"
                              fullWidth
                              margin="normal"
                              variant="outlined"
                              style={{ padding: "0 8px" }}
                            />
                          </CardContent>
                        </Card>
                      );
                    } else if (text === TEMPLATE_ELEMENT.NUMBER) {
                      const listKey = `form-string-${index}`;
                      return (
                        <Card key={listKey} style={{ margin: "8px" }}>
                          <CardHeader
                            avatar={<LooksOne />}
                            title="Number"
                            titleTypographyProps={{ variant: "subtitle1" }}
                          />
                          <CardContent>
                            <TextField
                              placeholder="Title"
                              margin="normal"
                              style={{ marginLeft: "8px" }}
                            />
                            <TextField
                              type="number"
                              placeholder="DefaultNumber"
                              fullWidth
                              margin="normal"
                              variant="outlined"
                              style={{ padding: "0 8px" }}
                            />
                          </CardContent>
                        </Card>
                      );
                    } else if (text === TEMPLATE_ELEMENT.NESTED) {
                      const listKey = `form-nested-${index}`;
                      return (
                        <Card key={listKey} style={{ margin: "8px" }}>
                          <CardHeader
                            avatar={<NestedIcon />}
                            title="Nested"
                            titleTypographyProps={{ variant: "subtitle1" }}
                          />
                          <CardContent>
                            <TextField
                              placeholder="Title"
                              margin="normal"
                              style={{ marginLeft: "8px" }}
                            />
                            <Card
                              style={{
                                backgroundColor: "#EEE",
                              }}
                            >
                              <Grid
                                container
                                alignItems="center"
                                justify="center"
                                style={{ height: "100px" }}
                              >
                                <Typography
                                  variant="subtitle2"
                                  color="textSecondary"
                                >
                                  Drag and Drop Area
                                </Typography>
                              </Grid>
                            </Card>
                          </CardContent>
                        </Card>
                      );
                    }
                  })}
                </FormGroup>
              </FormControl>
            </CardContent>
          </Card>
          <Box pt={4}>
            <Copyright />
          </Box>
        </Container>
      </main>
    </div>
  );
};

export default Main;
