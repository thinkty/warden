import React from 'react';
import { Item, Record } from './Item';

type Props = {
  interval: number;
  url: string;
  itemLen: string;
} & typeof defaultProps;

const defaultProps = {
  interval: 10000,
  url: '/data',
  itemLen: '150px',
};

export const Container = (props: Props): JSX.Element => {
  const [items, setItems] = React.useState<JSX.Element[]>([]);

  // Setup the fetch sequence on component mount
  React.useEffect(() => {
    const fetchRecords = (): void => {
      fetch(props.url)
      .then(response => response.json())
      .then(data => {
        const records = Array.from<Record>(data);
        const tempItems = records.map((record: Record) => <Item record={record}/>);

        if (tempItems.length === 0) {
          setItems([
            <h1>
              Empty...
            </h1>
          ]);
        } else {
          setItems([...tempItems]);
        }
      })
      .catch(err => {
        // TODO: handle error
        console.error(err);
      });
    }

    fetchRecords();
    const interval: NodeJS.Timeout = setInterval((): void => {
      fetchRecords();
    }, props.interval);

    return () => clearInterval(interval)
  }, []);

  return (
    <div
      style={{
        margin: '10px',
        display: 'grid',
        maxHeight: '95vh',
        width: '100vw',
        overflow: 'auto',
        gap: '10px',
        gridAutoRows: props.itemLen,
        gridTemplateColumns: `repeat(auto-fit, ${props.itemLen})`,
        justifyContent: 'center',
      }}
    >
      {
        items
      }
    </div>
  );
}
Container.defaultProps = defaultProps;
