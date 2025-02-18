import React from 'react';

const BackupCodes = ({recovery_codes = [], onGenerateNewCode}) => {
    const handleDownload = () => {
        const blob = new Blob([recovery_codes.map(el => el.code).join('\n')], {type: 'text/plain'});
        const url = URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        a.download = 'backup_codes.txt';
        a.click();
        URL.revokeObjectURL(url);
    };

    return (
        <div className="mt-6">
            {Array.isArray(recovery_codes) && (
                <>

                    <h4 className="text-lg font-semibold text-gray-200">Backup Codes</h4>

                    <ul className="mt-4 flex flex-wrap gap-2">
                        {recovery_codes?.map((item, index) => (
                            <li
                                key={index}
                                className="bg-gray-700 p-2 rounded-lg font-mono text-gray-100"
                            >
                                {item.code}
                            </li>
                        ))}
                    </ul>


                    <div className="flex items-center justify-start gap-x-2 mt-6">

                        <button
                            onClick={onGenerateNewCode}
                            className="btn btn-success text-white px-4 py-2 rounded-lg"
                        >
                            Generate New Code.
                        </button>

                        <button
                            onClick={handleDownload}
                            className="btn btn-primary text-white px-4 py-2 rounded-lg"
                        >
                            Download
                        </button>
                    </div>

                </>
            )}
        </div>
    );
};

export default BackupCodes;
