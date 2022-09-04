import React from "react";
import cn from "classnames";
import styles from "./index.module.scss";
import Text from "components/text";

export function TextInput({
  label = "",
  name,
  register,
  className,
  placeholder,
  errors,
  type = "text",
  ...props
}) {
  return (
    <div className={cn(styles.input, className)}>
      <label htmlFor={name}>{label ?? ""}</label>
      <input
        type={type}
        name={name}
        id={name}
        placeholder={placeholder}
        ref={register}
        className={cn({ error: errors })}
        {...props}
      />
      {errors && (
        <Text className="" tag="span" theme="small">
          {errors.message}
        </Text>
      )}
    </div>
  );
}
