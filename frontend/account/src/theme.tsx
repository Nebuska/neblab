import {createTheme} from "@mui/material/styles";
import {forwardRef} from "react";
import {
    type LinkProps as RouterLinkProps,
    Link as RouterLink,
} from 'react-router-dom';

const theme = createTheme({
    components: {
        MuiLink: {
            defaultProps: {
                component: forwardRef<
                    HTMLAnchorElement,
                    Omit<RouterLinkProps, 'to'> & { href: RouterLinkProps['to'] }
                >((props, ref) => {
                    const { href, ...other } = props;

                    return (<RouterLink ref={ref} to={href} {...other} />);
                }),
            }
        }
    }
})
export default theme;