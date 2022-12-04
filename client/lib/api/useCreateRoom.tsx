import { useMutation } from "@tanstack/react-query";
import api from "./api";

export interface RoomCreateRequest {
  name: string;
  password: string;
}

export interface RoomCreateResponse {
  id: string;
  name: string;
}

const useCreateRoom = () => {
  return useMutation(({ name, password }: RoomCreateRequest) =>
    api.post<RoomCreateRequest, RoomCreateResponse>(`/v1/room`, {
      name,
      password,
    })
  );
};

export default useCreateRoom;
