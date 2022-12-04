import { useMutation } from "@tanstack/react-query";
import api from "./api";

export interface RoomJoinRequest {
  name: string;
  password: string;
}

export interface RoomJoinResponse {
  id: string;
  name: string;
}

const useJoinRoom = () => {
  return useMutation(({ name, password }: RoomJoinRequest) =>
    api.post<RoomJoinRequest, RoomJoinResponse>(`/v1/room/join`, {
      name,
      password,
    })
  );
};

export default useJoinRoom;
