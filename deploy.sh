

echo "Deploying backend."
vercel deploy --prod


echo "Deploying frontend."
cd frontend
npm run build

vercel deploy --prod