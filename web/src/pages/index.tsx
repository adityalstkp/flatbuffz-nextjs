import type { GetServerSidePropsContext } from "next";
import FlatBufferUsers from "@/components/Home/FlatBufferUsers";
import JSONUsers from "@/components/Home/JSONUsers";

export default function Home(props: { userCount: number }) {
  const { userCount } = props;
  return (
    <main
      className={`flex min-h-screen flex-col items-center justify-between p-24`}
    >
      <FlatBufferUsers count={userCount} />
      <JSONUsers count={userCount} />
    </main>
  );
}

export async function getServerSideProps({ query }: GetServerSidePropsContext) {
  const userCount = Number(query.users_count) || 10;

  return {
    props: { userCount: userCount },
  };
}
