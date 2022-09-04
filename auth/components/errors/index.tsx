import React from "react";
import Text from "components/text";
import styles from "components/text-input/index.module.scss";

export function ErrorMsg({ messages = [], msg = "" }) {
  return (
    <div className={styles.error}>
      <Text className={styles.errorHead} tag="h5" theme="regular">
        {msg}
      </Text>
      {messages && (
        <ul>
          {messages.map((item, index) => (
            <li key={index}>
              <Text className="" tag="p" theme="small">
                {`* ${item}`}
              </Text>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
