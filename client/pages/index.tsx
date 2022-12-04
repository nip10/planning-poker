import { Center } from "@mantine/core";
import type { NextPage } from "next";
import RoomForm from "../components/RoomForm";

const Home: NextPage = () => {
  return (
    <Center>
      <RoomForm />
    </Center>
  );
};

export default Home;
