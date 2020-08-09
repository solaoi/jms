import { FC } from "react";
import { useDrag } from "react-dnd";
import { DRAG_EVENT_TYPE } from "./TemplateElementList";
import { ListItem, ListItemIcon, ListItemText } from "@material-ui/core";
import { EditAttributes, Create, LooksOne } from "@material-ui/icons";
import NestedIcon from "~/components/atoms/NestedIcon";

export const TEMPLATE_ELEMENT = {
  BOOLEAN: "boolean",
  STRING: "string",
  NUMBER: "number",
  NESTED: "nested",
} as const;

export type TemplateElementType = {
  type: string;
  text: string;
};

const TemplateElement: FC<{ text: string }> = ({ text }) => {
  const [, drag] = useDrag({
    item: {
      type: DRAG_EVENT_TYPE.TEMPLATE_ELEMENT,
      text,
    } as TemplateElementType,
  });

  return (
    <ListItem button ref={drag}>
      <ListItemIcon>
        {text === TEMPLATE_ELEMENT.BOOLEAN && <EditAttributes />}
        {text === TEMPLATE_ELEMENT.STRING && <Create />}
        {text === TEMPLATE_ELEMENT.NUMBER && <LooksOne />}
        {text === TEMPLATE_ELEMENT.NESTED && <NestedIcon />}
      </ListItemIcon>
      <ListItemText primary={text} />
    </ListItem>
  );
};

export default TemplateElement;
