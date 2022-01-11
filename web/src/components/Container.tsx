import React from 'react';
import { Item, SensorRecord } from './Item';

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

  // Setting up fetching records. Currently, the implementation is very crude
  // and it is just fetching all the records everytime which is very inefficient.
  // However, the optimization of fetching will come in the later versions.
  // Check the project board on Github on how to improve the fetch sequence.
  React.useEffect(() => {
    const fetchRecords = (): void => {
      fetch(props.url)
      .then(response => response.json())
      .then(data => {
        const records = Array.from<SensorRecord>(data);

        // Organize the records into groups by beacon and name
        const recordGroups = new Map<string, SensorRecord[]>();
        for (let i = 0; i < records.length; i++) {
          const record: SensorRecord = records[i];
          const key: string = record.Beacon + "/" + record.Name;

          // If the group already exists, update map. If not, create new entry
          if (recordGroups.has(key)) {
            const group = recordGroups.get(key);
            if (group) {
              group.push(record)
            } else {
              recordGroups.set(key, [record]);
            }
          } else {
            recordGroups.set(key, [record]);
          }
        }

        const tempItems: JSX.Element[] = [];
        for (const [_, group] of recordGroups.entries()) {
          tempItems.push(<Item records={group} />);
        }

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

    // Setup the fetch interval on component mount
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
      { items }
    </div>
  );
}
Container.defaultProps = defaultProps;
