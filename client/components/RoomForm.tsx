import {
  Box,
  TextInput,
  Checkbox,
  Group,
  Button,
  useMantineTheme,
  Flex,
} from "@mantine/core";
import { useForm } from "@mantine/form";
import { useRouter } from "next/router";
import useCreateRoom from "../lib/api/useCreateRoom";
import useJoinRoom from "../lib/api/useJoinRoom";

const RoomForm: React.FC = () => {
  const theme = useMantineTheme();
  const form = useForm({
    initialValues: {
      name: "",
      password: "",
    },
  });
  const createRoomMutation = useCreateRoom();
  const joinRoomMutation = useJoinRoom();
  const router = useRouter();

  const onCreateRoom = () => {
    console.log(form.values);
    createRoomMutation.mutate(form.values, {
      onSuccess: (data) => {
        console.log(data);
        router.push(`/room/${data.id}`);
      },
    });
  };

  const onJoinRoom = () => {
    console.log(form.values);
    joinRoomMutation.mutate(form.values, {
      onSuccess: (data) => {
        console.log(data);
        router.push(`/room/${data.id}`);
      },
    });
  };

  return (
    <Box
      sx={{
        width: 400,
        borderWidth: 1,
        borderColor: theme.colors.gray[2],
        borderStyle: "solid",
        borderRadius: theme.radius.sm,
        padding: theme.spacing.md,
        backgroundColor: theme.colors.gray[2],
      }}
      mx="auto"
    >
      <TextInput
        withAsterisk
        label="Name"
        placeholder="my room"
        {...form.getInputProps("name")}
      />

      <TextInput
        withAsterisk
        label="Password"
        placeholder="hunter2"
        {...form.getInputProps("password")}
      />

      <Group position="right" mt="md">
        <Button variant="outline" onClick={onJoinRoom}>
          Join
        </Button>
        <Button onClick={onCreateRoom}>Create</Button>
      </Group>
    </Box>
  );
};

export default RoomForm;
