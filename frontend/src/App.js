import React, { useState } from 'react'

import styles from './app.module.css';

function App() {
  const [kw, setKw] = useState("")
  const [imgURL, setImgURL] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    let url = new URL(window.location.href);
    url.href = "http://localhost:8080/";
    url.searchParams.set("kw", kw);
    url.searchParams.set("with_dots", "false");
    setImgURL(url.toString());
  };

  return (
    <>
      <form onSubmit={handleSubmit} className={styles.form}>
        <label htmlFor="kw">Keyword:</label>
        <input onInput={(e) => {setKw(e.target.value)}} className={styles.input} type="text" name="kw" placeholder="some secret string"/>
        <input className={styles.input} type="submit"/>
      </form>
      {!!imgURL && (
        // eslint-disable-next-line jsx-a11y/img-redundant-alt
        <img className={styles.img} src={imgURL} alt="Encoded image"/>
      )}
    </>
  );
}

export default App;
