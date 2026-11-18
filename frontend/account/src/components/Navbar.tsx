import {AppBar, Button, IconButton, Menu, MenuItem, Toolbar, Typography} from "@mui/material"
import MenuIcon from '@mui/icons-material/Menu';
import {Link} from "react-router-dom";
import * as React from "react";
import {useAuth} from "../contexts/AuthContext.tsx";

export default function Navbar() {
    const [user, signIn, signOut] = useAuth();

    const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
    const open = Boolean(anchorEl);
    const handleProfileButton = (event: React.MouseEvent<HTMLButtonElement>) => {
        setAnchorEl(event.currentTarget);
    };
    const handleClose = () => {
        setAnchorEl(null);
    };

    return (
        <>
            <AppBar>
                <Toolbar>
                    <IconButton
                    size="large"
                    edge="start"
                    color="inherit"
                    aria-label="open drawer"
                    sx={{
                        mr:2,
                        display: {
                            xs: "flex",
                            md: "none",
                        },
                    }}>
                        <MenuIcon />
                    </IconButton>
                    <Typography variant="h6" component={Link} sx={{ flexGrow: 1 }} to={"/"}>
                            Neblab
                    </Typography>

                    { user == null && (
                        <>
                            <Button color="inherit" component={Link} to="/login">
                                Login
                            </Button>
                            <Button color="inherit" component={Link} to="/register">
                                Register
                            </Button>
                        </>
                    )}
                    { user != null && (
                      <>
                          <Button
                              id="profile-button"
                              color="inherit"
                              aria-controls={open ? 'profile-menu' : undefined}
                              aria-haspopup="true"
                              aria-expanded={open ? 'true' : undefined}
                              onClick={handleProfileButton}
                          >
                              Profile
                          </Button>
                          <Menu
                              id="profile-menu"
                              anchorEl={anchorEl}
                              open={open}
                              onClose={handleClose}>

                              <MenuItem> Profile </MenuItem>
                              <MenuItem> Settings </MenuItem>
                              <MenuItem> Logout </MenuItem>
                          </Menu>
                      </>
                    )}

                </Toolbar>
            </AppBar>
            <Toolbar/>
        </>
    )
}