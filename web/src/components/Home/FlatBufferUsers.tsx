import { ByteBuffer } from "flatbuffers";
import {
  type FC,
  useState,
  Profiler,
  type ProfilerOnRenderCallback,
} from "react";
import { Users } from "../../../flatbuffer_gen/users";

interface Props {
  count: number;
}

const FlatBufferUsers: FC<Props> = (props: Props) => {
  const { count } = props;
  const [users, setUsers] = useState<Users | null>(null);

  const handleFetchAndRender = async () => {
    try {
      const resp = await fetch(
        `http://localhost:8000/flatbuff?users_count=${count}`
      );
      console.time("flatbuff-deserialize");
      const buffResp = await resp.arrayBuffer();
      const buff = new ByteBuffer(new Uint8Array(buffResp));
      setUsers(Users.getRootAsUsers(buff));
      console.timeEnd("flatbuff-deserialize");
    } catch (e) {
      console.error("[fetch][FlatBuffer]", e);
    }
  };

  const handleProfilerOnRender: ProfilerOnRenderCallback = (
    id,
    _,
    actualDuration,
    __,
    ___,
    ____
  ) => {
    console.log(`${id} took ${actualDuration}ms to render.`);
  };

  const renderUsers = (userBuff: Users) => {
    let usersList = [];
    let i = 0;
    const listLen = userBuff.listLength();
    console.time("flatbuff-while-loop");
    while (i < listLen) {
      const user = userBuff.list(i);
      const name = user?.name();
      usersList.push(
        <div key={`idx-${i}`} style={{ marginBottom: 5 }}>
          <p>{name}</p>
        </div>
      );
      i++;
    }
    console.timeEnd("flatbuff-while-loop");
    return usersList;
  };

  return (
    <section>
      <h1>FlatBuffer</h1>
      <p>Count: {count}</p>
      <button
        style={{ border: "1px solid red" }}
        onClick={handleFetchAndRender}
      >
        Render
      </button>
      <Profiler id="flatbuffer-render" onRender={handleProfilerOnRender}>
        <div style={{ height: 200, overflowY: "auto" }}>
          {users && renderUsers(users)}
        </div>
      </Profiler>
    </section>
  );
};

export default FlatBufferUsers;
