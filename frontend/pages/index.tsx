import { useState, useEffect } from "react";
import Calendar from "react-calendar";
import "react-calendar/dist/Calendar.css";

const Home = () => {
  const [date, setDate] = useState<Date | null>(new Date());
  const [memo, setMemo] = useState("");
  const [memos, setMemos] = useState<{ [key: string]: string }>({});

  useEffect(() => {
    if (date) {
      // 選択した日付のメモを取得
      fetch(`/api/memos?date=${date.toISOString().split("T")[0]}`)
        .then((response) => response.json())
        .then((data) => setMemo(data.memo || ""));
    }
  }, [date]);

  const handleSave = () => {
    if (date) {
      fetch("/api/memos", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ date: date.toISOString().split("T")[0], memo }),
      }).then(() => {
        setMemos({ ...memos, [date.toISOString().split("T")[0]]: memo });
      });
    }
  };

  return (
    <div>
      <Calendar onChange={(value) => setDate(value as Date)} value={date} />
      <div>
        <h2>{date?.toDateString()}</h2>
        <textarea value={memo} onChange={(e) => setMemo(e.target.value)} />
        <button onClick={handleSave}>保存</button>
      </div>
    </div>
  );
};

export default Home;
