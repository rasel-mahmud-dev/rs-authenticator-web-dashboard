import React, {useState} from 'react';
// Import Swiper React components
import {Swiper, SwiperSlide} from 'swiper/react';

// Import Swiper styles
import 'swiper/css';
import 'swiper/css/free-mode';
import 'swiper/css/navigation';
import 'swiper/css/thumbs';
import {FreeMode, Navigation, Thumbs} from 'swiper/modules';


const AppPreviewGallery = ({images = []}) => {
    const [thumbsSwiper, setThumbsSwiper] = useState(null);

    const isMobile = window.innerWidth < 576

    return (
        <>
            <Swiper
                style={{
                    '--swiper-navigation-color': '#fff',
                    '--swiper-pagination-color': '#fff',
                }}
                centeredSlides={true}
                navigation={true}
                thumbs={{swiper: thumbsSwiper}}
                modules={[FreeMode, Navigation, Thumbs]}
                className="mySwiper2"
            >
                {images?.map(el => (
                    <SwiperSlide key={el}>
                        <div className="android-frame w-[280px] md:w-[300px] mx-auto ">
                            <img alt={el} src={el}/>
                        </div>
                    </SwiperSlide>
                ))}
            </Swiper>
            <Swiper
                onSwiper={setThumbsSwiper}
                spaceBetween={5}
                slidesPerView={isMobile ? 5: 10}
                freeMode={true}
                watchSlidesProgress={true}
                modules={[FreeMode, Navigation, Thumbs]}
                className="mySwiper mt-4"
            >
                {images?.map(el => (
                    <SwiperSlide key={el}>
                        <div className="android-frame frame-xs ">
                            <div className="h-20">
                                <img alt={el} src={el}/>
                            </div>
                        </div>
                    </SwiperSlide>
                ))}
            </Swiper>
        </>
    );
};

export default AppPreviewGallery;