import React, { FC } from "react";
import TemplateElement, { TEMPLATE_ELEMENT } from "./TemplateElement";

export const DRAG_EVENT_TYPE = {
  TEMPLATE_ELEMENT: "template-element",
} as const;

const TEMPLATE_ELEMENT_LIST = [
  TEMPLATE_ELEMENT.BOOLEAN,
  TEMPLATE_ELEMENT.STRING,
  TEMPLATE_ELEMENT.NUMBER,
  TEMPLATE_ELEMENT.NESTED,
];

const TemplateElementList: FC = () => {
  return (
    <>
      {TEMPLATE_ELEMENT_LIST.map((text, index) => (
        <TemplateElement key={`template-element-${index}`} text={text} />
      ))}
    </>
  );
};

export default TemplateElementList;
