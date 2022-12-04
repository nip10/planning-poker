import React, { useState } from "react";
import {
  Burger,
  Header as MantineHeader,
  MediaQuery,
  useMantineTheme,
  Text,
} from "@mantine/core";

const Header: React.FC = () => {
  const theme = useMantineTheme();
  const [opened, setOpened] = useState(false);
  return (
    <MantineHeader height={{ base: 50, md: 70 }} p="md">
      <div style={{ display: "flex", alignItems: "center", height: "100%" }}>
        <MediaQuery largerThan="sm" styles={{ display: "none" }}>
          <Burger
            opened={opened}
            onClick={() => setOpened((o) => !o)}
            size="sm"
            color={theme.colors.gray[6]}
            mr="xl"
          />
        </MediaQuery>

        <Text>ACME</Text>
      </div>
    </MantineHeader>
  );
};

export default Header;
