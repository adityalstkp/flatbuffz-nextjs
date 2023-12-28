import {
  type FC,
  useState,
  Profiler,
  type ProfilerOnRenderCallback,
} from "react";

interface Props {
  count: number;
}

interface Geo {
  lat: string;
  long: string;
}

interface Address {
  geo: Geo;
}

interface User {
  name: string;
  address: Address;
}

const renderUsers = (users: User[]) => {
  let usersList = [];
  let i = 0;
  console.time("json-while-loop");
  while (i < users.length) {
    const name = users[i].name;
    usersList.push(
      <div key={`idx-${i}`} style={{ marginBottom: 5 }}>
        <p>{name}</p>
      </div>
    );
    i++;
  }
  console.timeEnd("json-while-loop");
  return usersList;
};

const JSONUsers: FC<Props> = (props: Props) => {
  const { count } = props;
  const [users, setUsers] = useState<User[]>([]);

  const handleFetchAndRender = async () => {
    try {
      const resp = await fetch(
        `http://localhost:8000/json?users_count=${count}`
      );
      console.time("json-deserialize");
      const userList = await resp.json();
      setUsers(userList);
      console.timeEnd("json-deserialize");
    } catch (e) {
      console.error("[fetch][JSON]", e);
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

  return (
    <section>
      <h1>JSON</h1>
      <p>Count: {count}</p>
      <button
        style={{ border: "1px solid red" }}
        onClick={handleFetchAndRender}
      >
        Render
      </button>
      <Profiler id="json-render" onRender={handleProfilerOnRender}>
        <div style={{ height: 200, overflowY: "auto" }}>
          {users.length > 0 && renderUsers(users)}
        </div>
      </Profiler>
    </section>
  );
};

export default JSONUsers;
