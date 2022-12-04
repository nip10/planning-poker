import { useRouter } from "next/router";

const Room = () => {
  const router = useRouter();
  const { id } = router.query;

  return <div>Room {id}</div>;
};

export default Room;
