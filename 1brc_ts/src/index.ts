import { compute } from 'src/compute';

async function main() {
  // const filePath = '../1brc/measurements.txt';
  const filePath = './data/weather_stations.csv';

  try {
    const results = await compute(filePath);
    console.log(results);

  } catch (e) {
    console.error(e);
  }
}

main()
