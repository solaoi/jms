import { Typography, Link } from "@material-ui/core";
import { FC } from "react";

const Copyright: FC = () => {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {"Copyright © "}
      <Link color="inherit" href="https://aota.blog/">
        aota.blog
      </Link>{" "}
      {new Date().getFullYear()}
      {"."}
    </Typography>
  );
};

export default Copyright;
