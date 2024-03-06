// Tokyo;35.6897
// Jakarta;-6.1750
// Delhi;28.6100
// Guangzhou;23.1300
// Mumbai;19.0761
// Manila;14.5958
// Shanghai;31.1667
// São Paulo;-23.5500
// Seoul;37.5600


// {Abha=5.0/18.0/27.4, Abidjan=15.7/26.0/34.1, Abéché=12.1/29.4/35.6, Accra=14.7/26.4/33.1, Addis Ababa=2.1/16.0/24.3, Adelaide=4.1/17.3/29.7, ...}
// result format: min, mean, max

interface result {
    city: string;
    min: number;
    mean: number;
    max: number;
}

async function compute(): Promise<Map<string, result>> {
    return new Map<string, result>();
}

async function main() {
    const results = compute();
    console.log(results);
}

main()